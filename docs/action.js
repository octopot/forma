(function ($) {
    'use strict';
    const location = new URL(window.location.href),
        body = $('body'),
        success = 'success',
        failure = 'failure';
    $('.example form').each(function (i, node) {
        const value = location.searchParams.get(node.id);
        if (value) {
            var msg, title;
            switch (value) {
                case success:
                    title = 'Form "' + node.title + '" was processed successfully!';
                    msg = $('<div class="msg alert alert-success alert-dismissible fade show"/>');
                    msg.append($('<h4 class="alert-heading"/>').text(title));
                    msg.append($('<p>Aww yeah, you did the right thing!</p>'));
                    break;
                case failure:
                    title = 'Form "' + node.title + '" was processed unsuccessfully';
                    msg = $('<div class="msg alert alert-danger alert-dismissible fade show"/>');
                    msg.append($('<h4 class="alert-heading"/>').text(title));
                    msg.append($('<p>Oops! But this also happens 😉</p>'));
                    break;
                default:
                    return;
            }
            msg.append($('<hr>'));
            msg.append($('<p class="mb-0">As you can see it was very simple! 🤗</p>'));
            msg.append(
                $('<button type="button" class="close" data-dismiss="alert"/>')
                    .append($('<span>&times;</span>'))
            );
            body.append(msg);
            setTimeout(function () { msg.alert('close'); }, 5000)
        }
    });
}(window.jQuery));
