import Data.Sequence (findIndexL, fromList)
import Data.List (nub)
import Data.Maybe (fromJust)

windowed size ls =
    case ls of
        [] -> []
        x:xs ->
            if length ls >= size then
                take size ls : windowed size xs
            else windowed size xs

unique ls = nub ls == ls

firstUnique = findIndexL unique . fromList

solve size = map (fmap (size +) . firstUnique . windowed size)

main = do
    input <- lines <$> readFile "input6.txt"
    print (solve 4 input)
    print (solve 14 input)
