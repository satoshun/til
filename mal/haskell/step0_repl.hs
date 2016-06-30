import Readline (readline)

loop = do
    line <- readline "user> "
    case line of
        Just str -> do
            putStrLn line
            loop

main = do
    loop
