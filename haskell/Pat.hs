pat :: [Int] -> Int
pat (0:1:2:xs) = 0
pat (0:1:xs) = 1
pat (0:[]) = 3
pat (0:xs) = 4
pat _ = 6

main = do
    print (pat [1, 2])
    print (pat [0, 5])
    print (pat [0])
    print (pat [0, 1, 2])
