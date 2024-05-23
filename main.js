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
 // Initialize nodes
let node1 = new Node(1);
let node2 = new Node(2);
let node3 = new Node(3);

// Link the nodes
node1.next = node2;
node2.next = node3;

// Create linked list and set the head
let linkedList = new  LinkList();
linkedList.head = node1;

//previous = null null->currentNode1
function reverses(node){
    
    //null<- node1->next->node2->node3->null
    //previous 
    //current 
    //next

    let current = node;
    let prev = null;
  
    while(current !=null){
       let temp = current.next; 
       current.next = prev; 
       prev = current;
       current = temp;
    
    }
    
    return prev;
  
}

linkedList.head = reverses(linkedList.head);
 
// Function to display the list
function displayList(head) {
    let current = head;
    while (current !== null) {
        console.log(current.data);
        current = current.next;
    }
}

displayList(linkedList.head)

