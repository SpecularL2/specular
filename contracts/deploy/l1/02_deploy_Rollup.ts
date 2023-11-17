import { exec } from "child_process";
import util from "node:util";
import path from "node:path";
import { HardhatRuntimeEnvironment } from "hardhat/types";
import { DeployFunction } from "hardhat-deploy/types";

import { deployUUPSProxiedContract, getProxyName } from "../utils";

const CLIENT_SBIN_DIR = `${__dirname}/../../../sbin`;

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  // Calculate initial VM hash
  const execPromise = util.promisify(exec);
  let initialVMHash = "";
  let stdout = "";
  // TODO: read genesis hash file directly
  const cmd = `sh ${path.join(CLIENT_SBIN_DIR, "read_genesis_hash.sh")}`;
  try {
      ({ stdout } = await execPromise(cmd));
  } catch (error) {
    throw Error(`could not export genesis hash: ${error}`);
  }
  try {
    initialVMHash = (JSON.parse(stdout).hash || "") as string;
  } catch (err) {
    throw Error(`could not parse ${stdout}: ${err}`);
  }
  if (!initialVMHash) {
     throw Error(`hash not found\n${stdout}`);
  }
  console.log("initial VM hash:", initialVMHash);

  const { deployments, getNamedAccounts } = hre;
  const { sequencer, validator, deployer } = await getNamedAccounts();
  const sequencerInboxProxyAddress = (
    await deployments.get(getProxyName("SequencerInbox"))
  ).address;
  const verifierProxyAddress = (await deployments.get(getProxyName("Verifier")))
    .address;

  const args = [
    sequencer, // address _vault
    sequencerInboxProxyAddress, // address _sequencerInbox
    verifierProxyAddress, // address _verifier
    5, // uint256 _confirmationPeriod
    0, // uint256 _challengePeriod
    0, // uint256 _minimumAssertionPeriod
    0, // uint256 _baseStakeAmount
    0, // uint256 _initialAssertionID
    0, // uint256 _initialInboxSize
    initialVMHash, // bytes32 _initialVMhash
    [sequencer, validator], // address[] calldata _validators
  ];

  await deployUUPSProxiedContract(hre, deployer, "Rollup", args);
};

export default func;
func.tags = ["Rollup", "L1", "Stage0"];
