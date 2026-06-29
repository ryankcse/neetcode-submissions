class Solution {
    /**
     * @param {string[]} strs
     * @return {string[][]}
     */
    groupAnagrams(strs) {
        let groups = {};

        strs.forEach((str) => {
            let occ = Array(26).fill(0);
            // iterate through each character
            for(const c of str) {
                occ[c.charCodeAt(0) - 'a'.charCodeAt(0)]++
            }
            // need to encode the map as a string or something...
            let encodedKey = occ.join("#")

            if(!groups[encodedKey]) {
                groups[encodedKey] = [];
            } 
            groups[encodedKey].push(str)
        })
        return Object.values(groups);
    }
}
