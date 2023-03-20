// The simplest double pointer
function twoSum (nums, target) {
  for (let i = 0; i < nums.length; i++) {
    for (let j = i + 1; j < nums.length; j++) {
      if (nums[i] + nums[j] === target) {
        return [i, j]
      }
    }
  }
  return null
}

// use hash table
function twoSumHash (nums, target) {
  const keyMap = {}
  for (const index in nums) {
    // 配对值
    const partner = target - nums[index]
    if (keyMap[partner] && index !== keyMap[partner]) {
      const t = val => parseInt(val)
      return [t(keyMap[partner]), t(index)]
    }
    keyMap[nums[index]] = index
  }
  return null
}

function main () {
  // test
  const testList = [
    { nums: [2, 7, 11, 15], target: 9, result: [0, 1] },
    { nums: [3, 2, 4], target: 6, result: [1, 2] },
    { nums: [3, 3], target: 6, result: [0, 1] }
  ]

  // log
  for (let i = 0; i < testList.length; i++) {
    const testTask = testList[i]
    const { nums, target, result } = testTask
    const res = twoSum(nums, target)
    const resHash = twoSumHash(nums, target)
    console.log(`| Task ${i}:`)
    console.log(`-------------------------------`)
    console.log(`|  nums = `, nums)
    console.log(`|  target = `, target)
    console.log(`|  result = `, result)
    console.log(`|  res = `, res)
    console.log(`|  resHash = `, resHash)
    console.log(`-------------------------------`)
  }
}

main()