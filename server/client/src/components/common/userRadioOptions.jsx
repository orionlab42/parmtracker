import React from "react";


const UserRadioOptions = ({users, onUserShare}) => {
    return (
        <div className="users-checkbox">
            {users.length && users.map(user =>
                <div key={user.user_id}>
                    <input id={user.user_id} className="check-with-label" type="checkbox" name="answer" onClick={() => onUserShare(user)}/>
                    <label className="checkbox label-for-check" htmlFor={user.user_id}>{user.user_name}</label>
                </div>)}
        </div>
    );
};

export default UserRadioOptions;