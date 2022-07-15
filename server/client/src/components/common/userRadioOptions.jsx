import React, {useEffect, useState} from "react";
import {getUsersOfNote} from "../../services/noteService";


const UserRadioOptions = ({users, note, onUserShare}) => {
    const [usersToShareStatus, setUsersToShareStatus] = useState([]);
    const [allShared, setAllShared] = useState(false);

    useEffect(() => {
        getNoteUsersToShare().then();
    }, [users]);

    console.log("users", users)
    console.log("usersToShare", usersToShareStatus)
    const getNoteUsersToShare = async () => {
        const { data: usersOfNote } = await getUsersOfNote(note.note_id);
        if (usersOfNote != null) {
            let usersShareStatus = [];
            let usersNotShared = 0;
            for (let i = 0; i < users.length; i++) {
                let notFound = true
                for (let j = 0; j < usersOfNote.length; j++) {
                    if (users[i].user_id === usersOfNote[j]) {
                        notFound = false;
                        usersShareStatus.push({user_id: users[i].user_id, user_name: users[i].user_name, shared: true});
                        break;
                    }
                }
                if (notFound) {
                    usersShareStatus.push({user_id: users[i].user_id, user_name: users[i].user_name, shared: false});
                    usersNotShared ++;
                }
            }
            setUsersToShareStatus(usersShareStatus);
            if (usersNotShared === 0) {
                setAllShared(true);
            }
        }
    }

    return (
        <div className="users-checkbox">
            {usersToShareStatus && usersToShareStatus.map(user => {
                if (user.shared === false)
                    return <div key={user.user_id}>
                        <input id={user.user_id}
                               className="check-with-label"
                               type="checkbox"
                               name="answer"
                               onClick={() => onUserShare(user.user_id)}/>
                        <label className="checkbox label-for-check" htmlFor={user.user_id}>{user.user_name}</label>
                    </div>
            }
                )}
        {allShared && <p>Already shared.</p>}
        </div>
    );
};

export default UserRadioOptions;