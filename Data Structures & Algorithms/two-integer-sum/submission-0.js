class Solution {
    /**
     * @param {number[]} nums
     * @param {number} target
     * @return {number[]}
     */
    twoSum(nums, target) {
        // we just need to know if an element exists
        let occ = new Map();
        for(let i = 0; i < nums.length; i++) {
            if(!occ.get(nums[i])) {
                occ.set(nums[i], true)
            }
        }

        for(let i = 0; i < nums.length; i++) {
            // if a complement value exists
            if(occ.get(target - nums[i])) {
                console.log(`i=${i}, target-nums[i]=${target-nums[i]}`)
                // figure out the index
                for(let j = i+1; j < nums.length; j++) {
                    if(nums[j] == target - nums[i]) {
                        return [i, j]
                    }
                }
            }
        }
        
    }
}
