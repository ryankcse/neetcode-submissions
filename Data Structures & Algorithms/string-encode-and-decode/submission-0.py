class Solution:

    def encode(self, strs: List[str]) -> str:
        encodedStr = ""
        for s in strs:
            encodedStr += (str(len(s)) + "#" + s)
        return encodedStr

    def decode(self, s: str) -> List[str]:
        head = 0
        decodedStrings = []
        while head < len(s):
            strLenAsStr = ""
            while s[head] != "#":
                strLenAsStr += s[head]
                head += 1
            head += 1 # skip the #
            strLen = int(strLenAsStr)
            decodedString = ""
            for i in range(strLen):
                decodedString += s[head]
                head += 1
            decodedStrings.append(decodedString)
        return decodedStrings

