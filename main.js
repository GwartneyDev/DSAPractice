 
class ListNode {
  constructor(val, next = null) {
    this.val = val;
    this.next = next;
  }
}

  
  function GetHead(head){
      let prev = null;     // Initialize prev to null
      let current = head;  // Start with head as the current node
     
      while(current !== null){
        let nextTemp =  current.next; // next temp currently is 2
        current.next = prev; // set next to previous first time around null
        prev = current; // set prev to 1 isntead of null 
        current =  nextTemp ;// current head moves to 2
        
        // nextTemp = head.Next = 2
        // head.next = prev = Null
        //prev = head prev now = 1
        //head = nextTemp head now is 2 and its next value is 1 

         //nextTemp = head.next = 3
        // head.next = prev = 1
        // prev = head prev now =2
        //head = nextTempp head is now 3 and its next is 2 

       //this repeats it should then be 3->2->1
      }
    return prev;
  }

 
 function printList(head) {
  let current = head;
  while (current !== null) {
    console.log(current.val);
    current = current.next;
  }
  console.log("null");
}

// Create a sample linked list: 1 -> 2 -> 3
let head = new ListNode(1);
head.next = new ListNode(2);
head.next.next = new ListNode(3);

// Reverse the list
let get = GetHead(head);

// Print the reversed list
 
printList(get);
