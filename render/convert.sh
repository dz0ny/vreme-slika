mkdir temp
for fullfile in svgs/*.svg; do
    filename=$(basename -- "$fullfile")
    extension="${filename##*.}"
    filename="${filename%.*}"
    /Applications/Inkscape.app/Contents/MacOS/inkscape -z -o /Users/dz0ny/vreme-fb-post/render/temp/$filename.png -w 512 -h 512 $fullfile
done
