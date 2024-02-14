const {greetUser} = require('./script.js');
// const {greetUser} = module.exports('./javascript/script');

describe("Function", () => {
    test('Greeting the user with a message', () => {screenLeft
        expect(greetUser()).toBe("Hello there, friend!");
    });
})

// const { TextEncoder } = require('text-encoding');
// const { JSDOM } = require('jsdom');

// const { fireEvent } = require('@testing-library/dom');
// require('@testing-library/jest-dom');

// const { document } = (new JSDOM('')).window;
// global.document = document;
// global.window = document.defaultView;
// const { addTask } = require('./script.js');

// // // Mock the DOM environment
// // beforeAll(() => {
// //     const div = document.createElement('div');
// //     div.id = 'tasks';
// //     document.body.appendChild(div);

// //     const input = document.createElement('input');
// //     input.id = 'newtask';
// //     div.appendChild(input);

// //     const pushButton = document.createElement('button');
// //     pushButton.id = 'push';
// //     div.appendChild(pushButton);

// //     // Mock the onclick event for the push button
// //     pushButton.onclick = jest.fn();
// // });

// // describe("addTask function", () => {
// //     test("Adds a task to the DOM", () => {
// //         addTask("Test Task");
// //         // Assert that the onclick event for the push button is called
// //         expect(document.querySelector('#push').onclick).toHaveBeenCalled();
// //     });
// // });


// // Import the addTask function from script.js
// // const { addTask } = require('./script.js');

// // Mock the window.onload event
// beforeEach(() => {
//     document.body.innerHTML = `
//         <div id="tasks"></div>
//         <div id="newtask">
//             <input type="text" id="taskInput" />
//             <button id="push">Push</button>
//         </div>
//     `;
    
//     // Mock the onclick event for the push button
//     document.querySelector('#push').onclick = jest.fn();
    
//     // Simulate window.onload event by calling addTask
//     addTask();
// });

// describe("addTask function", () => {
//     test("Adds a task to the DOM", () => {
//         // Check if task is added to the DOM
//         const taskElement = document.querySelector(".task");
//         expect(taskElement).not.toBeNull();
        
//         // Check task name
//         const taskNameElement = document.querySelector("#taskname");
//         expect(taskNameElement.textContent).toBe("Test Task");
//     });

//     test("Does not add task when input is empty", () => {
//         // Simulate clicking the button without entering task name
//         const pushButton = document.querySelector("#push");
//         pushButton.click();

//         // Check if task is not added to the DOM
//         expect(document.querySelector(".task")).toBeNull();
//     });

//     test("Deletes task when delete button is clicked", () => {
//         // Simulate adding a task
//         document.querySelector("#taskInput").value = "Test Task";
//         const pushButton = document.querySelector("#push");
//         pushButton.click();

//         // Simulate clicking the delete button
//         const deleteButton = document.querySelector(".delete");
//         deleteButton.click();

//         // Check if task is deleted from the DOM
//         expect(document.querySelector(".task")).toBeNull();
//     });
// });
