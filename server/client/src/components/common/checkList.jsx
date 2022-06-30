import React, {useState, useEffect} from "react";
import CheckListForm from "./checkListForm";
import CheckListItems from "./checkListItems";

const CheckList = ({ items, handleUpdateCheckList, handleDeleteCheckList }) => {
    const [checkList, setCheckList] = useState([]);

    useEffect(() => {
        setCheckList(items);
    }, [items]);

    const addItem = item => {
        // console.log("Adding item", item);
        if (item.text.trim().length > 0) {
            let newCheckList = checkList;
            newCheckList.list = [...checkList.list, item];
            setCheckList(newCheckList);
            handleUpdateCheckList(checkList);
        }
        console.log("CheckList: Adding item ", checkList);
    };


    const completeItem = (id) => {
        let newCheckList = checkList;
        newCheckList.list = checkList.list.map(item => {
            if (item.id === id) {
                item.isComplete = !item.isComplete;
            }
            return item
        });
        setCheckList(newCheckList);
        handleUpdateCheckList(checkList);
    };

    const deleteItem = (id) => {
        let newCheckList = checkList;
        newCheckList.list = checkList.list.filter(item => item.id !== id);
        setCheckList(newCheckList);
        handleUpdateCheckList(checkList);
    };

    const updateItem = (id, newValue, isComplete) => {
        let newCheckList = checkList;
        newCheckList.list = checkList.list.map(item => {
            if (item.id === id) {
                item.text = newValue.text;
                item.isComplete = isComplete;
            }
            return item
        });
        setCheckList(newCheckList);
        handleUpdateCheckList(checkList);
    };

    return (
       <div className="checklist-container">
           <div className="checklist-top">
                <CheckListForm onSubmit={addItem}/>
                <button className="button is-link is-light  mdi mdi-format-title" onClick={() => handleDeleteCheckList(items.id)}/>
           </div>
           <CheckListItems
               items={items.list}
               handleCompleteItem={completeItem}
               handleDeleteItem={deleteItem}
               handleUpdateItem={updateItem}/>
           <button className="button is-link is-light  mdi mdi-trash-can-outline" onClick={() => handleDeleteCheckList(items.id)}/>
       </div>
    );
}

export default CheckList;