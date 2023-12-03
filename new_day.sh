set -e
DAY=$1

# if DAY is not a number, exit
if ! [[ $DAY =~ ^[0-9]+$ ]]; then
    echo "Usage: $0 <day>"
    exit 1
fi


[ ${#DAY} -eq 1 ] && DAY="0$DAY"
[ $DAY -gt 25 ] && echo "Day must be between 1 and 25" && exit 1
[ $DAY -lt 1 ] && echo "Day must be between 1 and 25" && exit 1

DAY_FOLDER="day$DAY"

[ -d $DAY_FOLDER ] && echo "Folder for day $DAY already exists" && exit 1

mkdir $DAY_FOLDER
cp day00/day00_1.go $DAY_FOLDER/$DAY_FOLDER\_1.go
cp day00/day00_2.go $DAY_FOLDER/$DAY_FOLDER\_2.go

touch data/full/$DAY_FOLDER.txt
touch data/test/$DAY_FOLDER.txt


