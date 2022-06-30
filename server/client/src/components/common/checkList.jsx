import React, {useState} from "react";
import CheckListForm from "./checkListForm";
import CheckListItems from "./checkListItems";

const CheckList = ({ items, handleAddItem, handleCompleteItem, handleDeleteItem, handleUpdateItem}) => {
    return (
       <div className="checklist-container">
           <CheckListForm onSubmit={handleAddItem}/>
           <CheckListItems
               items={items}
               handleCompleteItem={handleCompleteItem}
               handleDeleteItem={handleDeleteItem}
               handleUpdateItem={handleUpdateItem}/>
       </div>
    );
}

export default CheckList;