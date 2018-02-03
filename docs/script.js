(function ($) {
    'use strict';
    const
        location = new URL(window.location.href),
        messages = $('#messages'), tpl = $('#message'),
        success = 'success',
        failure = 'failure';
    var hashMap = {};

    function showMessage(id, value) {
        const node = $('#' + id);
        var type, title, desc, message;
        switch (value) {
            case success:
                type = 'alert-success';
                title = '"' + node.attr('title') + '" form was processed successfully!';
                desc = 'Aww yeah, you did the right thing!';
                break;
            case failure:
                type = 'alert-danger';
                title = '"' + node.attr('title') + '" form processed unsuccessfully';
                desc = 'Oops! But this also happens 😉';
                break;
            default:
                return;
        }
        message = $(tpl.html()
            .replace('{{ type }}', type)
            .replace('{{ title }}', title)
            .replace('{{ desc }}', desc));
        messages.append(message);
        setTimeout(function () { message.alert('close'); }, 4000 + 1000 * Math.random());
    }

    function showMessages() {
        for (var id in hashMap) {
            if (hashMap.hasOwnProperty(id)) { showMessage(id, hashMap[id]); }
        }
        hashMap = {}
    }

    $('.example form').each(function (i, node) {
        const value = location.searchParams.get(node.id);
        if (value) { hashMap[node.id] = value }
    });
    showMessages();

    $('.clipboard .btn-alert').on('click', function (e) {
        e.preventDefault();
        $(this).parent().prev().find('form[id]').each(function (i, node) {
            hashMap[node.id] = Math.round(Math.random()) ? success : failure;
        });
        showMessages();
    });
}(window.jQuery));
