#! /bin/bash

goto()
{
	case "$1" in
		work)
			
			open ~/"OneDrive - De Lage Landen International B.V"
			;;
		downloads)
			open ~/Downloads
			;;
		desktop)
			open ~/Desktop
			;;
		documents)
			open ~/Documents
			;;
		nope)
			state="exit"
			;;
		*)
			echo "Sorry, I don't know that space"
			;;
		esac
}

if [ $# -ne 0 ]
then 
	goto "$1"
	exit
fi

state="welcome"

while : 
do
	case "$state" in
	welcome)
		echo "Hi Simin..."
		read -p 'which space do you want to go? ' space
		state="go"
		;;
	go) 
		goto "$space"
		if [ ! "$state" = "exit" ]
		then
			read -p 'anywhere else? ' space
		fi
		;;
	exit)
		echo "See you Simin"
		break
		;;
	*)
		;;
	esac
done

exit