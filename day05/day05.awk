#!/usr/bin/gawk -f

BEGIN { 
    FIELDWIDTHS = "4 4 4 4 4 4 4 4 4"
    max_height = 0
}
/\[/ {
    for (i = 1; i <= NF; i++) {
        if (substr($(i), 1, 3) != "   ") {
            cols[i][NR] = substr($(i), 1, 3)
        }
    }
    max_height++
}

/^$/ {
    FPAT = "[0-9]+" 

    # reversing each column so its highest index is the top of the stack
    PROCINFO["sorted_in"] = "@ind_num_desc"
    for (col in cols) {
        i = 1
        for (elem in cols[col]) {
            stacks[col][i] = cols[col][elem]
            stacks2[col][i] = cols[col][elem]
            i++
        }
    }
    printf("Starting state: \n")
    show_crates(stacks)
    printf("--------------------------------------------------------\n")
    instruction_count = 0
}

/move/ {
    instruction_count++
    crates_to_move = $1
    source = $2
    dest = $3

    for (i = 1; i <= crates_to_move; i++) {
        source_top = length(stacks[source])
        dest_top = length(stacks[dest]) + 1
        if (source_top in stacks[source]) {
            stacks[dest][dest_top] = stacks[source][source_top]
            delete stacks[source][source_top]
        }
    }

    for (i = 1; i <= crates_to_move; i++) {
        source_top = length(stacks2[source])
        move_stack[i] = stacks2[source][source_top]
        delete stacks2[source][source_top]
    }
    for (i = crates_to_move; i > 0; i--) {
        dest_top = length(stacks2[dest]) + 1
        stacks2[dest][dest_top] = move_stack[i]
    }

}

END {
    printf("part 1:\n")
    show_crates(stacks)
    printf("--------------------------------------------------------\n")
    printf("part 2:\n")
    show_crates(stacks2)
}

function show_crates(columns,    j) {
    old_sorted_in = PROCINFO["sorted_in"]
    PROCINFO["sorted_in"] = "@ind_num_asc"
    max = 0
    for (col in columns) {
        if (length(columns[col]) > max) {
            max = length(columns[col])
        }
    }
    max_height = max
    for (j = max_height; j > 0; j--) {
        for (col in columns) {
            if (j in columns[col]) {
                printf("%s ",columns[col][j])
            } else {
                printf("    ")
            }
        }
        printf("\n")
    }
    for (col in columns) {
        printf(" %d  ", col)
    }
    printf("\n")
    PROCINFO["sorted_in"] = old_sorted_in
}

