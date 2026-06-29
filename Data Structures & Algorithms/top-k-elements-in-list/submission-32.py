class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        from queue import PriorityQueue 
        output: List[int] = []
        q = PriorityQueue()
        
        occurrenceMap: Dict[str, int] = {}
        for i in nums:
            if not occurrenceMap.get(str(i)):
                occurrenceMap[str(i)] = 1
            else:
                occurrenceMap[str(i)] += 1
        
        for o in occurrenceMap:
            # note: priorityqueue orders items from lowest priority first.
            q.put((-occurrenceMap[o], o))

        # print(q)
        # while not q.empty():
        #     print(q.get()[1])
        # print(q.get())
        # print(type(q.get()))
        for i in range(k):
            output.append(int(q.get()[1]))

        return output