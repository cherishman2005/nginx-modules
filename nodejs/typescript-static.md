# TypeScript - Static

ES6 includes static members and so does TypeScript. The static members of a class are accessed using the class name and dot notation, without creating an object e.g. <ClassName>.<StaticMember>.

The static members can be defined by using the keyword static. Consider the following example of a class with static property.

Example: Static Property Copy
```
class Circle {
    static pi: number = 3.14;
}
```
The above Circle class includes a static property pi. This can be accessed using Circle.pi. TypeScript will generate the following JavaScript code for the above Circle class.
```
var Circle = /** @class */ (function () {
    function Circle() {
    }
    Circle.pi = 3.14;
    return Circle;
}());
```
The following example defines a class with static property and method and how to access it.

Example: Static Members Copy
```
class Circle {
    static pi: number = 3.14;
    
    static calculateArea(radius:number) {
        return this.pi * radius * radius;
    }
}
```
Circle.pi; // returns 3.14
Circle.calculateArea(5); // returns 78.5
The above Circle class includes a static property and a static method. Inside the static method calculateArea, the static property can be accessed using this keyword or using the class name Circle.pi.

Now, consider the following example with static and non-static members.

Example: Static and Non-static Members Copy
```
class Circle {
    static pi = 3.14;
    pi = 3;
}

Circle.pi; // returns 3.14

let circleObj = new Circle();
circleObj.pi; // returns 3
```
As you can see, static and non-static fields with the same name can exists without any error. The static field will be accessed using dot notation and the non-static field can be accessed using an object.

Let's see another example.

Example: Static and Non-static Members Copy
```
class Circle {
    static pi = 3.14;

    static calculateArea(radius:number) {
        return this.pi * radius * radius;
    }

    calculateCircumference(radius:number):number { 
        return 2 * Circle.pi * radius;
    }
}

Circle.calculateArea(5); // returns 78.5

let circleObj = new Circle();
circleObj.calculateCircumference(5) // returns 31.4000000
//circleObj.calculateArea(); <-- cannot call this
```
In the above example, the Circle class includes static method calculateArea and non-static method calculateCircumference. As you can see, the static field pi can be accessed in the static method using this.pi and in the non-static (instance) method using Circle.pi.

 Note:
The class or constructor cannot be static in TypeScript.

# link

- [https://www.tutorialsteacher.com/typescript/typescript-static](https://www.tutorialsteacher.com/typescript/typescript-static)