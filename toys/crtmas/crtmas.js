function crtmas(elem, imgData, srcWidth, srcHeight){

    var bgColor = 'hsla(0, 0%, 0%, 0.002)';

    var width = elem.width;
    var height = elem.height;

    var c = elem.getContext('2d');

    var grid = [];
    var data = imgData.data;

    var row = 0;
    var col = 0;

    for (let i = 0; i < data.length; i += 4){
        if (col >= srcWidth) {
            row++;
            col = 0
        }

        if (col == 0) {
            grid[row] = [];
        }

        var r = data[i];
        var g = data[i+1];
        var b = data[i+2];

        grid[row][col] = `rgba(${r}, ${g}, ${b}, 0.9)`;
        
        col++;
    }

    var state = {
        pointer: 0
    };

    var pixelWidth = (width / grid[0].length);
    var pixelHeight = (height / grid.length) * 0.9;


    (function draw(t){
        drawBg(c);

        var y = (height / grid.length) * state.pointer;
        
        var row = grid[state.pointer];

        for (var i = 0; i < row.length; i++){

                var x = (width / row.length) * i;
                

                c.save();
                c.beginPath();
                c.rect(x, y, pixelWidth, pixelHeight);
                c.fillStyle = row[i];
                c.fill();
                c.closePath();
                c.restore();
        }

        state.pointer++;
        if (state.pointer >= grid.length){
            state.pointer = 0;
        }

        requestAnimationFrame(draw);
    })(0);

    function drawBg(c){
        c.save();
        c.rect(0, 0, width, height)
        c.fillStyle = bgColor;
        c.fill();
        c.restore();
    }
}

