function copyContent(button) {
  // Get the description text from the data attribute
  var description = button.getAttribute('data-description');
  
  // Copy the description to the clipboard
  navigator.clipboard.writeText(description).then(() => {
    console.log('Content copied to clipboard');
  }).catch(err => {
    console.error('Failed to copy: ', err);
  });
}

