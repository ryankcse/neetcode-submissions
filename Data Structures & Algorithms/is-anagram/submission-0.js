class Solution {
    /**
     * @param {string} s
     * @param {string} t
     * @return {boolean}
     */
    isAnagram(s, t) {
        let occ = new Map();
        if(s.length !== t.length) {
            return false
        }
        for(let i = 0; i < s.length; i++) {
            
            occ.set(s.charAt(i), (occ.get(s.charAt(i)) || 0) + 1)
            occ.set(t.charAt(i), (occ.get(t.charAt(i)) || 0) - 1)
        }
        
        for(let [key, value] of occ) {
            if (value !== 0) {
                return false
            }
        }
        
        return true
    }
}
