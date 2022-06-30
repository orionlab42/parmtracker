import React, {useState} from "react";
import CheckListForm from "./checkListForm";

const CheckListItems = ({ items, handleCompleteItem, handleDeleteItem, handleUpdateItem }) => {
    const [edit, setEdit] = useState({
        id: null,
        text: '',
        isComplete: false
    });

    const submitUpdate = newValue => {
        handleUpdateItem(edit.id, newValue, edit.isComplete);
        // console.log("handleUpdate", edit);
        setEdit({
            id: null,
            text: '',
            isComplete: false});
    };

    if (edit.id) {
        return <CheckListForm edit={edit} onSubmit={submitUpdate}/>;
    }
    // console.log("Checklist in items", items);
    return (
        <div>
            {items.map(item => (
                <div key={item.id} className={item.isComplete ? 'checked checklist-item' : 'checklist-item'}>
                    <div key={item.id} onClick={() => handleCompleteItem(item.id)}>
                        <span>+</span> {item.text}
                    </div>
                    <div>
                        <button onClick={() => handleDeleteItem(item.id)}><span className="mdi mdi-trash-can-outline"/></button>
                        <button onClick={() => setEdit({id: item.id, text: item.text, isComplete: item.isComplete})}><span className="mdi mdi-circle-edit-outline"/></button>
                    </div>
                </div>
            ))}
        </div>
    );
}

export default CheckListItems;