def twoSum(self, nums: List[int], target: int) -> List[int]:
        # for index in length of input arr
        for i in range(len(nums)):
            # next element in range from 1 to end-remaining
            for j in range(i + 1, len(nums)):
                # if current + next equals target return indices of elements
                if nums[i] + nums[j] == target:
                    return [i,j]