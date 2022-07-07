import React, {useState, useEffect} from "react";
import CheckListForm from "./checkListForm";
import CheckListItems from "./checkListItems";

const CheckList = ({ items, handleUpdateCheckList, onDeleteCheckList }) => {
    const [idItem, setIdItem] = useState(1);
    const [checkList, setCheckList] = useState([]);
    const [titleOn, setTitleOn] = useState(false);

    useEffect(() => {
        setCheckList(items);
    }, [items]);

    const addItem = item => {
        if (item.text.trim().length > 0) {
            let newCheckList = checkList;
            newCheckList.list = [...checkList.list, item];
            setCheckList(newCheckList);
            handleUpdateCheckList(checkList);
        }
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

    const renderTitleInput = () => {
        setTitleOn(!titleOn);
    };

    const titleChange = (e) => {
        let newCheckList = checkList;
        newCheckList.title = e.target.value
        setCheckList(newCheckList);
        handleUpdateCheckList(checkList);
    }

    const title = (
        <input
            className="title-input edit"
            type="text"
            placeholder="Title here..."
            value={items.title}
            name="text"
            onChange={titleChange}
        />
    );

    return (
        <div className="note checklist">
            <div className="checklist-container">
                {!titleOn && <h4 className="note-title">{items.note_title}</h4>}
                {titleOn && title}
                <div className="checklist-top">
                    <CheckListForm onSubmit={addItem} newIdItem={idItem} increaseIdItem={setIdItem}/>
                    <button className="button is-link is-light  mdi mdi-format-title"
                            onClick={renderTitleInput}/>
                </div>
                <div className="checklist-main">
                    <div className="checklist-body">
                        {!items.note_empty && <CheckListItems
                            items={items.note_list}
                            handleCompleteItem={completeItem}
                            handleDeleteItem={deleteItem}
                            handleUpdateItem={updateItem}
                            newIdItem={idItem}
                            increaseIdItem={setIdItem}/>}
                        <div className="note-footer">
                            <small>{!items.note_empty ? "Last modified:" +  new Date(items.updated_at).toLocaleDateString("en-GB", {
                                hour: "2-digit",
                                minute:  "2-digit",
                            }) : ""}</small>
                        </div>
                    </div>
                    <button className="button is-link is-light  mdi mdi-trash-can-outline"
                            onClick={() => onDeleteCheckList(items.note_id)}/>
                </div>
            </div>
        </div>
    );
}

export default CheckList;