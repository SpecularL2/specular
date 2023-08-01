import { useState } from 'react';
import { ethers } from 'ethers';
import { getStorageKey, requestFundDeposit } from '../utils';
import type { PendingDeposit, MessageProof } from "../types";
import {
  IL2Portal__factory,
  L1Oracle__factory,
} from "../typechain-types";

import {
  CHIADO_NETWORK_ID,
  CHIADO_RPC_URL,
  SPECULAR_NETWORK_ID,
  SPECULAR_RPC_URL,
  L1PORTAL_ADDRESS,
  L2PORTAL_ADDRESS,
  L1ORACLE_ADDRESS,
  DEPOSIT_BALANCE_THRESHOLD,
} from "../constants";
import {
  NETWORKS
} from "../chains";
interface Data {
  status: string;
  error?: string;
  data?: string;
}

interface wallet {
    address: string;
    chainId: number;
    provider: any;
  }



const INITIAL_DATA: Data = { status: 'pending' };

async function generateDepositProof(
  deposit: PendingDeposit
): Promise<MessageProof> {
  if (deposit.proofL1BlockNumber === undefined) {
    throw new Error("proofL1BlockNumber is undefined");
  }
  let rawProof = undefined;
  while (rawProof === undefined) {
    try {
      rawProof = await (new ethers.providers.StaticJsonRpcProvider(CHIADO_RPC_URL)).send(
        "eth_getProof",
        [
          L1PORTAL_ADDRESS,
          [getStorageKey(deposit.depositHash)],
          ethers.utils.hexValue(deposit.proofL1BlockNumber),
        ]
      );
    } catch (e) {
      console.error(e);
    }
    await new Promise((resolve) => setTimeout(resolve, 1000));
  }
  return {
    accountProof: rawProof.accountProof,
    storageProof: rawProof.storageProof[0].proof,
  };
}
type SwitchChainFunction = (arg: string) => void;

function useFinalizeDeposit(switchChain: SwitchChainFunction) {
  const [data, setData] = useState<Data>(INITIAL_DATA);

  const finalizeDeposit = async (wallet: wallet, amount: ethers.BigNumberish, pendingDeposit:PendingDeposit): Promise<void> => {

    if (!wallet) {
      setData({ status: 'failed', error: "Wallet doesn't exist" });
      return;
    }
    const l2Provider = new ethers.providers.StaticJsonRpcProvider(SPECULAR_RPC_URL)
    const l2Balance  = await l2Provider.getBalance(wallet.address);

    setData({ status: 'loading' });
    console.log("FInalizing Bro with l2 banance"+l2Balance);
    const targetBalance = ethers.utils.parseEther(ethers.utils.formatUnits(l2Balance, NETWORKS[SPECULAR_NETWORK_ID].nativeCurrency.decimals));
    if (DEPOSIT_BALANCE_THRESHOLD.gt(targetBalance)) {
      console.log("Sending Request Bro");
      try {
        const txHash = await requestFundDeposit(pendingDeposit);
        console.log("Success Transaction :"+txHash);
        setData({ status: 'successful', data: txHash });
      } catch (e) {
        console.error(e);
      }
      return;
    }
    console.log("NotSending Request Bro");
    switchChain(SPECULAR_NETWORK_ID.toString());
    const provider = await wallet.provider
    const signer = await (provider as any).getSigner();
    const l2Portal = IL2Portal__factory.connect(
      L2PORTAL_ADDRESS,
      signer,
    );
    const l1Oracle = L1Oracle__factory.connect(
      L1ORACLE_ADDRESS,
      provider,
    );
    try {
      var latestBlockNumber = await l1Oracle.blockNumber();
      while(!latestBlockNumber.gte(pendingDeposit.l1BlockNumber)) {
        latestBlockNumber = await l1Oracle.blockNumber();
      }
      pendingDeposit.proofL1BlockNumber = latestBlockNumber.toNumber();
      const proof = await generateDepositProof(pendingDeposit);
      const tx = await l2Portal.finalizeDepositTransaction(
        pendingDeposit.depositTx,
        proof.accountProof,
        proof.storageProof
      );
      setData({ status: 'pending', data: tx.hash });
      await tx.wait();
      setData({ status: 'successful', data: tx.hash });
    } catch (e) {
      console.error(e);
    }
    switchChain(CHIADO_NETWORK_ID.toString());
  };

  return { finalizeDeposit, data };
}

export default useFinalizeDeposit;
