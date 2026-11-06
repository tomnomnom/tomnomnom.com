function waves(elem){

    var bgColor = 'hsl(0, 0%, 0%)';
    var baseFreq = 2;

    var width = elem.width;
    var height = elem.height;

    var c = elem.getContext('2d');

    // square approximates a square wave by adding the third, fifth,
    // and seventh harmonics at one third, one fifth and one seventh
    // of their amplitudes respectively.
    var square = function(n){
        return (
            Math.sin(n) +
            Math.sin(n*3)/3 +
            Math.sin(n*5)/5 +
            Math.sin(n*7)/7
        );
    };

    // This can actually draw more than one wave if you want it to.
    // Each wave has:
    //   frac - how far through the animation the wave is as a float (0 to 1)
    //   step - how far to progress the animation each frame
    //   tail - how many dots should be drawn
    //   dots - an array to contain which fractions to draw dots at
    //   fn   - the function to generate the wave, e.g y = fn(x)
    var waves = [
        {frac: 0, step: 0.002, tail: 100, dots: [], fn: square}
    ];

    // Main draw loop
    var draw = function(){

        // Draw the background layer to cover up the previous frame
        drawBg(c);

        // Draw all of the waves
        for (var i = 0; i < waves.length; i++){
            var w = waves[i];

            drawWave(c, w);

            // Update wave's state
            w.frac += w.step;
            if (w.frac > (1+w.step)){
                w.frac = 0;
            }

        }

        // Draw the next frame when the browser is ready
        requestAnimationFrame(draw);
    };

    function drawBg(c){
        c.save();
        c.rect(0, 0, width, height)
        c.fillStyle = bgColor;
        c.fill();
        c.restore();
    }

    function drawWave(c, w){
        c.save();

        // Translate the y axis to the half way point as our waves'
        // y coordinate will fall above and below zero
        c.translate(0, height/2);

        // The 'shadow' is used as a bloom effect, mostly to appease @CharlotteGore
        c.shadowBlur = 30;

        // w.dots is kind of a like a ring buffer containing all of the fractions
        // (essentially unscaled x values) at which we want to draw a dot.
        // The 'head' dot is the first in the array
        w.dots.unshift(w.frac);

        // To keep the tail the right length we need to pop anything extra off the end
        if (w.dots.length > w.tail){
            w.dots.pop();
        }

        // Draw all the dots
        for (var i = 0; i < w.dots.length; i++){
            // f is a fraction of how far we are through the animation
            // e.g. 0.5 is half way through the wave
            var f = w.dots[i];

            // Use the width of the canvas element to 'scale' the x value
            var x = f*width;

            // The y value is produced by the wave's own function, and then scaled
            // to half of 90% of the canvas height.
            var y = w.fn(f*2*Math.PI*baseFreq) * (height*0.9/2);

            // Use the f value to decide on a hue and the 'dot number' to
            // decide on an alpha amount
            var hue = f*360;
            var alpha = 1 - (i/w.dots.length);
            var color = 'hsla('+hue+', 80%, 50%, '+alpha+')';

            // Draw the dot
            c.beginPath();
            c.arc(x, y, 5, 0, 2 * Math.PI);
            c.shadowColor = color;
            c.fillStyle = color;
            c.fill();
            c.closePath();
        }

        c.restore();
    }

    // Trigger the animation
    draw(0);
}

