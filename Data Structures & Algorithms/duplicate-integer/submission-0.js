class Solution {
    /**
     * @param {number[]} nums
     * @return {boolean}
     */
    hasDuplicate(nums) {
        let map = new Map()
        for (let i = 0; i < nums.length; i++) {
            if (map.get(nums[i]) == true) {
                return true
            } 
            map.set(nums[i], true)
        }
        return false
    }
}
