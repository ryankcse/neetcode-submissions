from collections import defaultdict 
class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        '''
        
    Each row must contain the digits 1-9 without duplicates.
    Each column must contain the digits 1-9 without duplicates.
    Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without duplicates.

        '''

        # Each row must contain the digits 1-9 without duplicates.
        for row in board:
            freq = Counter(row)
            for key, value in freq.items():
                if key != '.' and value > 1:
                    return False 
            

        # Each column must contain the digits 1-9 without duplicates.
        for col_index in range(9):
            col = []
            for row in board:
                col.append(row[col_index])
            freq = Counter(col)
            for key, value in freq.items():
                if key != '.' and value > 1:
                    return False 

        # Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without duplicates.
        # rows and columns 1, 4, and 7 will be the top left corner of a 3x3 square
        for row_index in range(0, 9, 3):
            for col_index in range(0, 9, 3):
                print("Square indices: ", row_index, col_index)
                # top left corner = board[row_index][col_index]
                square = []
                # push back row_index +0,1,2 and col_index+0,1,2
                print("Square coords")
                for row_offset in range(3):
                    for col_offset in range(3):
                        square.append(board[row_index+row_offset][col_index+col_offset])
                        print(row_index+row_offset, col_index+col_offset)
                freq = Counter(square)
                print("Counter:")
                for key, value in freq.items():
                    print(key, value)
                    if key != '.' and value > 1:
                        return False 


        return True