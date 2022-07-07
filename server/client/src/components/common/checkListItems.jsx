import React, {useState} from "react";
import CheckListForm from "./checkListForm";

const CheckListItems = ({ allItems, handleCompleteItem, handleDeleteItem, handleUpdateItem }) => {
    const [edit, setEdit] = useState({
        item_id: null,
        item_text: '',
        item_is_complete: false
    });

    const submitUpdate = newValue => {
        handleUpdateItem(edit.item_id, newValue, edit.item_is_complete);
        setEdit({
            item_id: null,
            item_text: '',
            item_is_complete: false});
    };

    if (edit.item_id) {
        return <CheckListForm edit={edit} onSubmit={submitUpdate}/>;
    }

    return (
        <div className="checklist-item-container">
            {allItems && allItems.map(item => (
                <div key={item.item_id}
                     className={item.item_is_complete ? 'checked checklist-item' : 'checklist-item note-content'}>
                    <div key={item.item_id}
                         onClick={() => handleCompleteItem(item.item_id)}>
                        <span>+</span> {item.item_text}
                    </div>
                    <div>
                        <button className="item-button"
                                onClick={() => handleDeleteItem(item.item_id)}><span className="mdi mdi-trash-can-outline"/></button>
                        <button className="item-button"
                                onClick={() => setEdit({item_id: item.item_id, item_text: item.item_text, item_is_complete: item.item_is_complete})}><span className="mdi mdi-circle-edit-outline"/></button>
                    </div>
                </div>
            ))}
        </div>
    );
}

export default CheckListItems;