const {greetUser} = require('./script.js');
// const {greetUser} = module.exports('./javascript/script');

describe("Function", () => {
    test('Greeting the user with a message', () => {
        expect(greetUser()).toBe("Hello there, friend!");
    });
})