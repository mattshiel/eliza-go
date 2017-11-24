// Author: Matthew Shiel
$("#user-input-form").submit(
    function (event) {
        event.preventDefault();

        // Store input in a variable
        var question = $('#user-input').val().trim() // Remove whitespace

        // Clear the input box
        $('#user-input').val("");

        // If the user doesn't input anything return nothing
        if (!question) {
            return;
        }

        // Append the user's question to the list and clear input
        $("#output-area").append('<li>' + question + '</li>');

        // scroll to the bottom of the messages
        // adapted from https://stackoverflow.com/questions/11715646/scroll-automatically-to-the-bottom-of-the-page
        $('#output-area').animate({
            scrollTop: $('#output-area').prop("scrollHeight")
        }, 300);

        $.get('/user-input', {
                value: question
            })
            .done(function (data) {
                // Set a timeout to make Eliza seem like she's thinking
                // Adapted random number in a range from here https://stackoverflow.com/questions/1527803/generating-random-whole-numbers-in-javascript-in-a-specific-range
                setTimeout(() => {
                    // Add Eliza's answer to the list
                    addListItem(data);
                }, Math.floor(Math.random() * (4000 - 600 + 1)) + 600);
            })
    });

function addListItem(input) {
    // Add the input to the list
    $("#output-area").append('<li>' + input + '</li>')
    // Automatically scroll to the last message 
    // adapted from https://stackoverflow.com/questions/11715646/scroll-automatically-to-the-bottom-of-the-page
    $('#output-area').animate({
        scrollTop: $('#output-area').prop("scrollHeight")
    }, 300);
}