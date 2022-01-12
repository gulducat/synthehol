default: run play
# default: run plot play

run:
	go run . > out.txt

plot: # brew install gnuplot
	gnuplot -e 'plot "out.txt" with boxes' -p

play: # brew install ffmpeg
	ffplay out.bin -f f32le -ar 44100 -x 1600 -y 800 -showmode 1

clean:
	rm out.*
