'use strict';

var gulp = require('gulp');
var sass = require('gulp-sass');
var gulp = require('gulp');
var cleanCSS = require('gulp-clean-css');
var minify = require('gulp-minify');

gulp.task('default', ['sass:watch', 'js:watch']);

// SASS compilation
gulp.task('sass', function () {
    return gulp.src('./styles/**/*.scss')
        .pipe(sass().on('error', sass.logError))
        .pipe(cleanCSS({compatibility: 'ie8'}))
        .pipe(gulp.dest('./static/css'));
});

gulp.task('sass:watch', function () {
    gulp.watch('./styles/**/*.scss', ['sass']);
});


// JS minify
gulp.task('js', function () {
    return gulp.src('./scripts/**/*.js')
        .pipe(minify({
            noSource: true
        }))
        .pipe(gulp.dest('./static/js'));
});

gulp.task('js:watch', function () {
    gulp.watch('./scripts/**/*.js', ['js']);
});
