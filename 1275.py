def tictactoe(moves: list[list[int]]) -> str:
    A = list()
    B = list()

    for i in range(0, len(moves)):
        if i % 2 == 0:
            A.append(moves[i])
        else:
            B.append(moves[i])
    
    if [0, 0] in A and [1, 1] in A and [2, 2] in A:
        return('A')
    elif [0, 0] in A and [0, 1] in A and [0, 2] in A:
        return('A')
    elif [1, 0] in A and [1, 1] in A and [1, 2] in A:
        return('A')
    elif [2, 0] in A and [2, 1] in A and [2, 2] in A:
        return('A')
    elif [0, 2] in A and [2, 0] in A and [1, 1] in A:
        return('A')
    elif [0, 0] in A and [1, 0] in A and [2, 0] in A:
        return('A')
    elif [0, 1] in A and [1, 1] in A and [2, 1] in A:
        return('A')
    elif [0, 2] in A and [1, 2] in A and [2, 2] in A:
        return('A')

    if [0, 0] in B and [1, 1] in B and [2, 2] in B:
        return('B')
    elif [0, 0] in B and [0, 1] in B and [0, 2] in B:
        return('B')
    elif [1, 0] in B and [1, 1] in B and [1, 2] in B:
        return('B')
    elif [2, 0] in B and [2, 1] in B and [2, 2] in B:
        return('B')
    elif [0, 2] in B and [2, 0] in B and [1, 1] in B:
        return('B')
    elif [0, 0] in B and [1, 0] in B and [2, 0] in B:
        return('B')
    elif [0, 1] in B and [1, 1] in B and [2, 1] in B:
        return('B')
    elif [0, 2] in B and [1, 2] in B and [2, 2] in B:
        return('B')
    
    if len(moves) != 9:
        return "Pending"
    return "Draw"
    
print(tictactoe([[0,0],[1,1],[0,1],[0,2],[1,0],[2,0]]))