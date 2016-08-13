'use strict';

var gulp = require('gulp');
var sass = require('gulp-sass');
var gulp = require('gulp');
var cleanCSS = require('gulp-clean-css');

gulp.task('default', function() {
});

gulp.task('sass', function () {
    return gulp.src('./styles/**/*.scss')
        .pipe(sass().on('error', sass.logError))
        .pipe(cleanCSS({compatibility: 'ie8'}))
        .pipe(gulp.dest('./static/css'));
});

gulp.task('sass:watch', function () {
    gulp.watch('./styles/**/*.scss', ['sass']);
});
