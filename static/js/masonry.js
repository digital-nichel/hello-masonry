const $grid = $('.grid').imagesLoaded(function () {
    $grid.masonry({
        gutter: '.gutter-sizer',
        itemSelector: '.grid-item',
        columnWidth: '.grid-sizer',
        percentPosition: true,
    });
});
