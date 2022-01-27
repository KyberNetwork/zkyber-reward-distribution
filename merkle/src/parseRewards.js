const {MerkleTree} = require('merkletreejs');
const ethers = require('ethers');
const BN = ethers.BigNumber;
const keccak256 = require('keccak256');
const ethersKeccak = ethers.utils.keccak256;
const abiEncoder = ethers.utils.defaultAbiCoder;

module.exports.parseRewards = function (rewardInfo) {
  const cycle = rewardInfo.cycle;
  const userRewards = rewardInfo.userRewards;
  // verify addresses (check duplicates, invalid) and convert cumulative amounts to BN
  // with account as key
  const mappedTokensAmounts = verifyAddressAndConvertAmounts(userRewards);
  // structure data to include account field
  const treeElements = addAccountInMapping(mappedTokensAmounts);
  // hash tree elements to leaves
  const leaves = hashElements(treeElements, cycle);
  const tree = new MerkleTree(leaves, keccak256, {sort: true});
  const userRewardsWithProof = treeElements.reduce((memo, {account}, index) => {
    let tokens = mappedTokensAmounts[account].tokens
    let amounts = mappedTokensAmounts[account].amounts.map((amt) => amt.toHexString())

    memo[account] = {
      index,
      tokens,
      amounts,
      proof: tree.getHexProof(leaves[index]),
    };
    return memo;
  }, {});

  return {
    cycle: cycle,
    startTimestamp: rewardInfo.startTimestamp,
    endTimestamp: rewardInfo.endTimestamp,
    merkleRoot: tree.getHexRoot(),
    userRewards: userRewardsWithProof,
  };
};

function verifyAddressAndConvertAmounts(userRewards) {
  return Object.keys(userRewards).reduce((memo, account) => {
    if (!ethers.utils.isAddress(account)) {
      throw new Error(`Found invalid address: ${account}`);
    }
    const parsedAddress = ethers.utils.getAddress(account);
    if (memo[parsedAddress]) throw new Error(`Duplicate address: ${parsed}`);
    const parsedTokenAmounts = userRewards[account].amounts.map((amt) => BN.from(amt));
    memo[parsedAddress] = {
      tokens: userRewards[account].tokens,
      amounts: parsedTokenAmounts,
    };
    return memo;
  }, {});
}

function addAccountInMapping(mappedTokensAmounts) {
  return Object.keys(mappedTokensAmounts).map((account) => ({
    account,
    tokens: mappedTokensAmounts[account].tokens,
    amounts: mappedTokensAmounts[account].amounts,
  }));
}

function hashElements(treeElements, cycle) {
  return treeElements.map((element, index) =>
    ethersKeccak(
      abiEncoder.encode(
        ['uint256', 'uint256', 'address', 'address[]', 'uint256[]'],
        [cycle.toString(), index.toString(), element.account, element.tokens, element.amounts]
      )
    )
  );
}
