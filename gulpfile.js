/*jshint node:true */
var gulp = require('gulp');
var spawn = require('child_process').spawn;
var gutil = require('gulp-util');

/* Create helper functions */

// param 'e' can be passed as a callback to the other 
// functions if you want.
var build = function (e) {
    'use strict';
    console.log('Currently in directory ', __dirname);

    var path = __dirname.split('/');
    if (path[path.length - 1] !== 'slide') {
        console.log('Error: in directory ', __dirname, ' should be in \'slide\'');
        return;
    }

    // recompile www/ directory into files 
    execute('go-bindata', ['-nomemcopy', 'www/...'], function () {
        
        // build the slide binary into $GOPATH/bin
        execute('go', ['install'], function () {
            if (isFunction(e)) {
                // if you pass a function as 'e', then
                // run it as a callback
                e();
            }
        });
    });    

}

// callback expects no args and will only execute when 
// the process completes. If no callback is passed, then
// you can chain the execute(cmd, args) functions 
// asyncronously. Only exec. callback if exit code of 
// previous command is 0.
var execute = function (cmd, args, callback) {
    'use strict';
    gutil.log("Running >> ", cmd, args);
    
    // Finally execute your script below - here "ls -lA"
    var child = spawn(cmd, args, {cwd: process.cwd()}),
        stdout = '',
        stderr = '';

    child.stdout.setEncoding('utf8');
    child.stdout.on('data', function (data) {
        stdout += data;
        gutil.log(data);
    });

    child.stderr.setEncoding('utf8');
    child.stderr.on('data', function (data) {
        stderr += data;
        gutil.log(gutil.colors.red(data));
        gutil.beep();
    });

    child.on('close', function (code) {
        gutil.log("Command", cmd, args, "done with exit code", code);
        
        // execute callback.
        if (code === 0 && isFunction(callback)) {
            callback();
        }
    });
};

var isFunction = function (functionToCheck) {
    'use strict';
    var getType = {};
    return functionToCheck && getType.toString.call(functionToCheck) === '[object Function]';
};

/* Define the gulp tasks */

gulp.task('default', function () {
    'use strict';
    
    var globs = ['www/*.js', 'www/*.css', 'www/*.html', 'www/img/*.{png,jpg,jpeg}'];
    
    gulp.watch(globs, build);
});

gulp.task('build', function () {
    'use strict';
    
    build('');
});

gulp.task('test', function () {
    'use strict';
    
    build(function () {
        // assumes that $GOPATH/bin is in $PATH
        execute('slide', ['-img=example/pdf']);
    })
});
