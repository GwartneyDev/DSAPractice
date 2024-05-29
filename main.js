class Node{
    constructor(data){
        this.next = null
        this.data = data;
    } 

}



class LinkList {
    constructor(){
        this.head = null;
    }
}




let node1 = new Node(1);
let node2 = new Node(2);
let node3 = new Node(3);

// Link the nodes
node1.next = node2;
node2.next = node3;

//previous = null null->currentNode1
function reverses(node){
    
    //tail<- node1->next->node2->node3->head
    //previous 
    //current 
    //next

    while (current != null) {
        next = current.next; // Store the next node
        current.next = prev; // Reverse the current node's pointer
        prev = current;      // Move prev to current node
        current = next;      // Move to the next node
    }

    return prev; // prev is the new head of the reversed list
}


reverses(current)


class ArrayClass {
    constructor() {
        this.length = 0;
    }

    // Method to add an element to the end of the array
    push(element) {
        this[this.length] = element;
        this.length++;
        return this.length;
    }

    // Method to add an element to the beginning of the array
    unshift(element) {
        // Shift elements to the right
        for (let i = this.length; i > 0; i--) {
            this[i] = this[i - 1];
        }
        // Place the new element at the start
        this[0] = element;
        // Increment the length
        this.length++;
        return this.length;
    }

    // Method to get an element by index
    get(index) {
        return this[index];
    }
}

// Example usage
let myArray = new ArrayClass();
myArray.push(1);
myArray.push(2);
myArray.push(3);
console.log(myArray); // Outputs: ArrayClass { '0': 1, '1': 2, '2': 3, length: 3 }

myArray.unshift(4);
console.log(myArray); // Outputs: ArrayClass { '0': 4, '1': 1, '2': 2, '3': 3, length: 4 }
