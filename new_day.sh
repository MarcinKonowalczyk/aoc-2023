set -e
XX=$1

# if XX is not a number, exit
if ! [[ $XX =~ ^[0-9]+$ ]]; then
    echo "Usage: $0 <day>"
    exit 1
fi


[ ${#XX} -eq 1 ] && XX="0$XX"
[ $XX -gt 25 ] && echo "Day must be between 1 and 25" && exit 1
[ $XX -lt 1 ] && echo "Day must be between 1 and 25" && exit 1

DAYXX="day$XX"

[ -d $DAYXX ] && echo "Folder for day $XX already exists" && exit 1

mkdir $DAYXX
cp day00/day00_1.go $DAYXX/$DAYXX\_1.go
cp day00/day00_2.go $DAYXX/$DAYXX\_2.go

# The first line of each of those files says 'package day00'
# We need to change that to 'package dayXX'

sed -i '' "s/day00/$DAYXX/g" $DAYXX/$DAYXX\_1.go
sed -i '' "s/day00/$DAYXX/g" $DAYXX/$DAYXX\_2.go

touch data/full/$DAYXX.txt
touch data/test/$DAYXX.txt

echo "data for $DAYXX" > data/full/$DAYXX.txt
echo "data for $DAYXX" > data/test/$DAYXX.txt
