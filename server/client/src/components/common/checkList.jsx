import React, {useState} from "react";
import CheckListForm from "./checkListForm";
import CheckListItems from "./checkListItems";

const CheckList = () => {
    const [items, setItems] = useState([]);

    const addItem = item => {
        if (item.text.trim().length > 0) {
            const newItems = [...items, item];
            setItems(newItems);
        }
    };

    const completeItem = (id) => {
        const updatedItems = items.map(item => {
            if (item.id === id) {
                item.isComplete = !item.isComplete;
            }
            return item
            });
        setItems([...updatedItems]);
    };

    const deleteItem = (id) => {
        const updatedItems = items.filter(item => item.id !== id);
        setItems([...updatedItems]);
    };

    const updateItem = (id, newValue, isComplete) => {
        const updatedItems = items.map(item => {
            if (item.id === id) {
                item.text = newValue.text;
                item.isComplete = isComplete;
            }
            return item
        });
        setItems([...updatedItems]);
    };

    return (
       <div className="checklist-container">
           <CheckListForm onSubmit={addItem}/>
           <CheckListItems
               items={items}
               handleCompleteItem={completeItem}
               handleDeleteItem={deleteItem}
               handleUpdateItem={updateItem}/>
       </div>
    );
}

export default CheckList;