loop ls = do
    case ls of
        (a : as) -> do
            print a
            loop as
        [] -> do
            print "empty"

main = do
    let a = Just 100
    let c = Nothing
    case a of
        Just b -> do
            print "Just"
        Nothing -> do
            print "Nothing"

    let d = [a, c]
    loop d
