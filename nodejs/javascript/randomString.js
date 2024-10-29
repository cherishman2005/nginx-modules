function randomString({ length, charset }) {
  let result = '';
  const characters = charset === 'numeric' ? '0123456789' : 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789';
  
  for (let i = 0; i < length; i++) {
    result += characters.charAt(Math.floor(Math.random() * characters.length));
  }
  
  return result;
}

randomString({ length: 8, charset: 'numeric' })


randomString({ length: 8}).toLowerCase()
