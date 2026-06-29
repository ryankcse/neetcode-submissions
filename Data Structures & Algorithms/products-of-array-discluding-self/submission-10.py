class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        prefix = []
        suffix = nums[1:]
        

        '''
example input [1,2,4,6]
index = 0
prefix = []
suffix = [2,4,6]

1 * 2 * 4 * 6 = 48

index = 1
prefix = [1]
suffix = [4,6]
what happened here?
1. we added nums[index - 1] to prefix
2. we changed suffix to nums[2:]
1 * 4 * 6 = 24

        '''
        output : List[int] = []
        for index in range(len(nums)):
            prefixProduct = 1
            if len(prefix) > 0:
                for p in prefix:
                    prefixProduct *= p
            suffixProduct = 1
            if len(suffix) > 0:
                for s in suffix:
                    suffixProduct *= s
            print(index, prefixProduct, suffixProduct)
            output.append(prefixProduct * suffixProduct)

            prefix.append(nums[index])
            suffix = nums[index+2:]
            print(prefix, suffix)
        return output