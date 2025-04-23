wait_for_key() {
    while true; do
        read -sn 3 key </dev/tty
        # "$" to interpret escape character
        if [[ "$key" == $'\e[C' ]] || [[ "$key" == $'\e[D' ]] || [[ "$key" == $'\e[B' ]]; then
            break
        fi
    done
}

# Function to generate unique random numbers
generate_random_numbers() {
    local count=$1
    local -A unique_numbers=()  # Associative array to track unique numbers
    local numbers=()
    local random_number

    # Generate unique random numbers
    while [ ${#numbers[@]} -lt $count ]; do
        random_number=$((RANDOM % 2001 - 1000))  # Generate random number
        if [[ -z "${unique_numbers[$random_number]}" ]]; then
            unique_numbers[$random_number]=1  # Mark this number as used
            numbers+=("$random_number")      # Add to the list
        fi
    done

    echo "${numbers[@]}"
}

./build.sh

echo "Try to run ./push-swap"
./push-swap
wait_for_key

echo "Try to run ./push-swap '2 1 3 6 5 8' "
./push-swap "2 1 3 6 5 8"
echo -n "Instructions counter: "
./push-swap "2 1 3 6 5 8" | wc -l
wait_for_key

echo "Try to run ./push-swap '0 1 2 3 4 5' "
./push-swap "0 1 2 3 4 5"
wait_for_key

echo "Try to run ./push-swap '0 one 2 3' "
./push-swap "0 one 2 3"
wait_for_key

echo "Try to run ./push-swap '1 2 2 3' "
./push-swap "1 2 2 3"
wait_for_key

echo "Try to run ./push-swap '<5 random numbers>' with 5 random numbers instead of the tag."
random_numbers=$(generate_random_numbers 5)
./push-swap "$random_numbers"
echo -n "Instructions counter: "
./push-swap "$random_numbers" | wc -l
wait_for_key

echo "Try to run ./push-swap '<5 random numbers>' with 5 different random numbers instead of the tag."
random_numbers=$(generate_random_numbers 5)
./push-swap "$random_numbers"
echo -n "Instructions counter: "
./push-swap "$random_numbers" | wc -l
wait_for_key

echo "Try to run ./checker and input nothing."
./checker
wait_for_key

echo "Try to run ./checker '0 one 2 3' "
./checker "0 one 2 3"
wait_for_key

echo "Try to run echo -e 'sa\npb\nrrr\n' | ./checker '0 9 1 8 2 7 3 6 4 5'"
echo -e "sa\npb\nrrr\n" | ./checker "0 9 1 8 2 7 3 6 4 5"
wait_for_key

echo "Try to run echo -e 'pb\nra\npb\nra\nsa\nra\npa\npa\n' | ./checker '0 9 1 8 2'"
echo -e "pb\nra\npb\nra\nsa\nra\npa\npa\n" | ./checker "0 9 1 8 2"
wait_for_key

echo "Try to run ARG='4 67 3 87 23); ./push-swap ARG | ./checker ARG"
ARG=('4 67 3 87 23')
./push-swap "$ARG" | ./checker "$ARG"
wait_for_key

echo "Try to run ARG=('<100 random numbers>'); ./push-swap ARG with 100 random different numbers instead of the tag."
ARG=$(generate_random_numbers 100)
echo $ARG
./push-swap "$ARG"
echo -n "Instructions counter: "
./push-swap "$ARG" | wc -l
wait_for_key

echo "Try to run ARG=('<100 random numbers>'); ./push-swap ARG | ./checker ARG with 100 random different numbers instead of the tag."
ARG=$(generate_random_numbers 100)
./push-swap "$ARG"  | ./checker "$ARG"
wait_for_key

rm checker push-swap
