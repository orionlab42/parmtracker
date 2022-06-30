import React from "react";
import CheckList from "./common/checkList";

const CheckListsList = ({checkLists, handleUpdateCheckList, handleDeleteCheckList}) => {
    return (
        <div className="notes-list-container">
            <h1 className="title is-3 center-text">NotesList</h1>
            <div className="notes-list">
                {checkLists.map(checkList => <div className="note"><CheckList
                    key={checkList.id}
                    items={checkList}
                    handleUpdateCheckList={handleUpdateCheckList}
                    handleDeleteCheckList={handleDeleteCheckList}
                /></div>)}
            </div>
        </div>
    );
};

export default CheckListsList;