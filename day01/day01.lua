#!/usr/bin/lua

function solve()
    io.input("input")
    local calorie_counts = {}
    local calories = 0
    for line in io.lines() do
        if line ~= "" then
            local food_item = tonumber(line)
            calories = calories + food_item
        else
            table.insert(calorie_counts, calories)
            calories = 0
        end
    end
    print(calorie_counts[1])
    print(calorie_counts[1]+calorie_counts[2]+calorie_counts[3])
end

solve()
