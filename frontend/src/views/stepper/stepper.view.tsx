import { useState, useCallback } from 'react'
import useStepperStyles from './stepper.styles'
import useWallet from '../../hooks/use-wallet'
import useStep, { Step } from '../../hooks/use-stepper-data'
import Login from '../login/login.view'
import DepositForm from '../deposit-form/deposit-form.view'
import WithdrawForm from '../withdraw-form/withdraw-form.view'
import TxConfirm from '../tx-confirm/tx-confirm.view'
import useDeposit from '../../hooks/use-deposit'
import useFinalizeDeposit from '../../hooks/use-finalize-deposit'
import TxPending from '../tx-pending/tx-pending.view'
import TxFinalizeDeposit from '../tx-finalize-deposit/tx-finalize-deposit.view'
import TxOverview from '../tx-overview/tx-overview.view'
import NetworkError from '../network-error/network-error.view'
import { ethers } from 'ethers'
import DataLoader from '../data-loader/data-loader'
import {
  CHIADO_NETWORK_ID,
  SPECULAR_NETWORK_ID,
  L1PORTAL_ADDRESS,
  CHIADO_RPC_URL,
  L1ORACLE_ADDRESS,
  SPECULAR_RPC_URL,
} from "../../constants";
import {
  IL1Portal__factory,
  IL2Portal__factory,
  L1Oracle__factory,
  IRollup__factory,
  ISequencerInbox__factory,
} from "../../typechain-types";
import type { PendingDeposit } from "../../types";

function Stepper () {
  const classes = useStepperStyles()
  const { wallet, loadWallet, disconnectWallet, isMetamask, switchChain } = useWallet()
  const { step, switchStep } = useStep()
  const { deposit, data: depositData, resetData: resetDepositData } = useDeposit()
  const { finalizeDeposit, data: finalizeDepositData, resetData: resetFinalizeDepositData } = useFinalizeDeposit()
  const [amount, setAmount] = useState(ethers.BigNumber.from("0"));

  const initialPendingDeposit: PendingDeposit = {
    l1BlockNumber: 0,
    proofL1BlockNumber: undefined,
    depositHash: "",
    depositTx: {
      nonce:ethers.BigNumber.from("0"),
      sender:"",
      target: "",
      value:ethers.BigNumber.from("0"),
      gasLimit:ethers.BigNumber.from("0"),
      data: ""
    }
  };
  const [pendingDeposit, setPendingDeposit] = useState(initialPendingDeposit)
  const [isDepositReadyToFinalize, setIsDepositReadyToFinalize] = useState(false)



  const tabs = [
    { name: 'Deposit', step: Step.Deposit },
  ]
  tabs.push({name: 'Withdraw', step: Step.Withdraw })

  const [activeTab, setActiveTab] = useState(tabs[0].name)

  const selectTab = useCallback((tab: { name: string; step: Step }) => {
    if (activeTab === tab.name) return;
    setActiveTab(tab.name);
    switchStep(tab.step);
  }, [activeTab, switchStep]);


  if (wallet && !(wallet.chainId == CHIADO_NETWORK_ID || wallet.chainId == SPECULAR_NETWORK_ID) ){
    return (
      <div className={classes.stepper}>
        <NetworkError {...{ isMetamask, switchChain }} />
      </div>
    )
  }

  const l1Provider = new ethers.providers.StaticJsonRpcProvider(CHIADO_RPC_URL)
  const l2Provider = new ethers.providers.StaticJsonRpcProvider(SPECULAR_RPC_URL)
  const l1Portal = IL1Portal__factory.connect(
    L1PORTAL_ADDRESS,
    l1Provider
  );
  const l1Oracle = L1Oracle__factory.connect(
    L1ORACLE_ADDRESS,
    l2Provider
  );

  l1Portal.on(
    l1Portal.filters.DepositInitiated(),
    (nonce, sender, target, value, gasLimit, data, depositHash, event) => {
      if (event.transactionHash === depositData.data?.hash) {
        const newPendingDeposit: PendingDeposit = {
          l1BlockNumber: event.blockNumber,
          proofL1BlockNumber: undefined,
          depositHash: depositHash,
          depositTx: {
            nonce,
            sender,
            target,
            value,
            gasLimit,
            data,
          },
        }
        setPendingDeposit(newPendingDeposit);
      }
    }
  );

  l1Oracle.on(
    l1Oracle.filters.L1OracleValuesUpdated(),
    (blockNumber, stateRoot, event) => {
      setIsDepositReadyToFinalize(false);
      if (pendingDeposit === undefined) {
        return;
      }
      if (blockNumber.gte(pendingDeposit.l1BlockNumber)) {
        setIsDepositReadyToFinalize(true);
        pendingDeposit.proofL1BlockNumber = blockNumber.toNumber();
      }
    }
  );

  return (
    <div className={classes.container}>
      {![Step.Login, Step.Loading].includes(step) && (
        <div className={classes.tabs}>
          {tabs.map(tab =>
            <button
              key={tab.name}
              className={activeTab === tab.name ? classes.tabActive : classes.tab}
              onClick={() => selectTab(tab)}
              disabled={![Step.Withdraw, Step.Deposit].includes(step)}
            >
              <span className={classes.tabName}>{tab.name}</span>
            </button>
          )}
        </div>
      )}
      <div className={classes.stepper}>
        {(() => {
          switch (step) {
            case Step.Loading: {
              console.log("Loading attempted")
              return (
                <DataLoader
                onGoToNextStep={() => switchStep(Step.Deposit)}
                />
              )
            }
            case Step.Login: {
              console.log("Login attempted")
              return (
                <Login
                  wallet={wallet}
                  onLoadWallet={loadWallet}
                  onGoToNextStep={() => switchStep(Step.Deposit)}
                />
              )
            }
            case Step.Deposit: {
              console.log("Deposit tab")
              switchChain(CHIADO_NETWORK_ID.toString())
              console.log("Chain Id is: "+wallet.chainId)
              return (
                <DepositForm
                  wallet={wallet}
                  depositData={depositData}
                  onAmountChange={resetDepositData}
                  onSubmit={(fromAmount) => {
                    deposit(wallet, fromAmount)
                    setAmount(fromAmount)
                    switchStep(Step.ConfirmDeposit)
                  }}
                  onDisconnectWallet={disconnectWallet}
                />
              )
            }
            case Step.Withdraw: {
              console.log("Withdraw tab")
              switchChain(SPECULAR_NETWORK_ID.toString())
              console.log("Chain Id is: "+wallet.chainId)
              return (
                <WithdrawForm
                  wallet={wallet}
                  depositData={depositData}
                  onAmountChange={resetDepositData}
                  onSubmit={(fromAmount) => {
                    deposit(wallet, fromAmount)
                    switchStep(Step.ConfirmWithdraw)
                  }}
                  onDisconnectWallet={disconnectWallet}
                />
              )
            }
            case Step.ConfirmDeposit: {
              console.log("Tx Confirmed")
              return (
                <TxConfirm
                  wallet={wallet}
                  depositData={depositData}
                  onGoBack={() => switchStep(Step.Deposit)}
                  onGoToPendingStep={() => switchStep(Step.PendingDeposit)}
                />
              )
            }
            case Step.ConfirmWithdraw: {
              console.log("Tx Confirmed")
              return (
                <TxConfirm
                  wallet={wallet}
                  depositData={depositData}
                  onGoBack={() => switchStep(Step.Withdraw)}
                  onGoToPendingStep={() => switchStep(Step.PendingWithdraw)}
                />
              )
            }
            case Step.PendingDeposit: {
              console.log("PendingDeposit")
              return (
                <TxPending
                  wallet={wallet}
                  depositData={depositData}
                  onGoBack={() => switchStep(Step.Deposit)}
                  onGoToOverviewStep={() => switchStep(Step.FinalizeDeposit)}
                />
              )
            }
            case Step.PendingWithdraw: {
              console.log("PendingWithdraw")
              return (
                <TxPending
                  wallet={wallet}
                  depositData={depositData}
                  onGoBack={() => switchStep(Step.Withdraw)}
                  onGoToOverviewStep={() => switchStep(Step.FinalizeWithdrawl)}
                />
              )
            }
            case Step.FinalizeDeposit: {
              console.log("FinalizeDeposit")
              finalizeDeposit(wallet,amount,pendingDeposit)
              return (
                <TxFinalizeDeposit
                  wallet={wallet}
                  depositData={depositData}
                  finalizeDepositData={finalizeDepositData}
                  onGoBack={() => switchStep(Step.Deposit)}
                  onGoToOverviewStep={() => switchStep(Step.Overview)}
                />
              )
            }
            case Step.FinalizeWithdrawl: {
              console.log("FinalizeWithdrawl")
              return (
                <TxPending
                  wallet={wallet}
                  depositData={depositData}
                  onGoBack={() => switchStep(Step.Withdraw)}
                  onGoToOverviewStep={() => switchStep(Step.Overview)}
                />
              )
            }
            case Step.Overview: {
              console.log("Overview")
              return (
                <TxOverview
                  wallet={wallet}
                  depositData={depositData}
                  onDisconnectWallet={disconnectWallet}
                  isMetamask={isMetamask}
                />
              )
            }
            default: {
              return <></>
            }
          }
        })()}

      </div>
    </div>
  )
}

export default Stepper
