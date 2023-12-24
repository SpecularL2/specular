import fs from "fs";
import { ethers } from "ethers";
import hre from "hardhat";
import { parseFlag } from "./utils";

// TODO: consider moving to golang (ops).
async function main() {
  const baseConfigPath = parseFlag("--in");
  const configPath = parseFlag("--out");
  const deploymentsConfig = parseFlag("--deployments-config-path");
  const genesisPath = parseFlag("--genesis");
  const genesisHashPath = parseFlag("--genesis-hash-path");
  const deploymentsPath = parseFlag("--deployments", "./deployments/localhost");
  await generateConfigFile(
    configPath,
    baseConfigPath,
    genesisPath,
    genesisHashPath,
    deploymentsPath,
  );
  await generateContractAddresses(deploymentsConfig, deploymentsPath);
}

/**
 * Reads the L1 and L2 genesis block info from the specified deployment and
 * adds it to the base config file.
 * Outputs the new config file at `configPath`.
 */
export async function generateConfigFile(
  configPath: string,
  baseConfigPath: string,
  genesisPath: string,
  genesisHashPath: string,
  deploymentsPath: string,
) {
  // check the deployments dir - error out if it is not there
  const contract = "Proxy__Rollup";
  const deployment = JSON.parse(
    fs.readFileSync(`${deploymentsPath}/${contract}.json`, "utf-8"),
  );
  const inboxDeployment = JSON.parse(
    fs.readFileSync(`${deploymentsPath}/Proxy__SequencerInbox.json`, "utf-8"),
  );

  // extract L1 block hash and L1 block number from receipt
  const l1Number = deployment.receipt.blockNumber;
  const l1Hash = deployment.receipt.blockHash;

  const baseConfig = JSON.parse(fs.readFileSync(baseConfigPath, "utf-8"));
  // Parse genesis hash file.
  const l2Hash = JSON.parse(fs.readFileSync(genesisHashPath, "utf-8")).hash;
  // Set genesis L1 fields.
  baseConfig.genesis.l1.hash = l1Hash;
  baseConfig.genesis.l1.number = l1Number;
  // Set genesis L2 fields.
  baseConfig.genesis.l2.hash = l2Hash;
  const genesis = JSON.parse(fs.readFileSync(genesisPath, "utf-8"));
  baseConfig.genesis.l2.number = ethers.BigNumber.from(genesis.number).toNumber();
  baseConfig.genesis.l2_time = ethers.BigNumber.from(genesis.timestamp).toNumber();
  baseConfig.genesis.gasLimit = ethers.BigNumber.from(genesis.gasLimit).toNumber();
  // Set other fields.
  baseConfig.l2_chain_id = genesis.config.chainId;
  baseConfig.batch_inbox_address = inboxDeployment.address;
  baseConfig.rollup_address = deployment.address;

  // Write out new file.
  fs.writeFileSync(configPath, JSON.stringify(baseConfig, null, 2));
  console.log(`successfully wrote config to: ${configPath}`);
}

/**
 * Reads the L1 deployment and writes deployments address to the deployments env file
 */
export async function generateContractAddresses(
  deploymentsConfigPath: string,
  deploymentsPath: string,
) {
  // check the deployments dir - error out if it is not there
  const deploymentFiles = fs.readdirSync(deploymentsPath);
  let result = "";
  for (const deploymentFile of deploymentFiles) {
    if (
      deploymentFile.startsWith("Proxy__") &&
      deploymentFile.endsWith(".json")
    ) {
      const deployment = JSON.parse(
        fs.readFileSync(`${deploymentsPath}/${deploymentFile}`, "utf-8"),
      );
      let contractName = deploymentFile
        .replace(/^Proxy__/, "")
        .replace(/\.json$/, "");
      contractName = contractName
        .replace(/([a-z])([A-Z])/g, "$1_$2")
        .toUpperCase();
      result += `${contractName}_ADDR=${deployment.address}\n`;
    }
  }
  fs.writeFileSync(deploymentsConfigPath, result);
  console.log(`successfully wrote deployments to: ${deploymentsConfigPath}`);
}

if (!require.main!.loaded) {
  main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
  });
}
