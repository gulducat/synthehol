default: run play
# default: run plot play

run:
	go run . > out.txt

plot: # brew install gnuplot
	gnuplot -e 'plot "out.txt" with boxes' -p

play: play-2

play-%: # brew install ffmpeg
	ffplay out.bin -f f32le -ar 44100  -x 1600 -y 800 -loop 8  -autoexit -showmode $(*)

clean:
	rm out.*
