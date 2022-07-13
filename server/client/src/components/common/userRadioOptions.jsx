import React from "react";


const UserRadioOptions = ({users, onUserShare}) => {
    return (
        <div className="users-checkbox">
            {users.length && users.map(user =>
                <div>
                    <label key={user.user_id} className="checkbox label-for-check" htmlFor={user.user_id}>{user.user_name}</label>
                    <input id={user.user_id} className="check-with-label" type="checkbox" name="answer" onClick={() => onUserShare(user)}/>
                </div>)}
        </div>
    );
};

export default UserRadioOptions;