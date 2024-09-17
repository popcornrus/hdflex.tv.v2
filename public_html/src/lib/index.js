// place files you want to import through the `$lib` alias in this folder.

if (!Object.prototype.hasOwnProperty('IsMovieType')) {
    Object.defineProperty(Object.prototype, 'IsMovieType', {
        value: function () {
            if (!this) return false;

            const types = [1, 3];
            return types.includes(this.content_type);
        },
        writable: false
    });
}

if (!Object.prototype.hasOwnProperty('IsTvSeriesType')) {
    Object.defineProperty(Object.prototype, 'IsTvSeriesType', {
        value: function () {
            if (!this) return false;

            const types = [2, 4, 5];
            return types.includes(this.content_type);
        },
        writable: false
    });
}