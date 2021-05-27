import React, {Component} from "react";
import {getCategories, deleteCategory} from "../services/categoryService";
import {toast} from "react-toastify";
import Link from "react-router-dom/Link";
import CategoriesList from "./categoriesList";
import CategoryForm from "./categoryForm";
import {Route} from "react-router-dom";



class Categories extends Component {
    state = {
        categories: [],
        sortColumn: {path: 'id', order: 'asc'},
        addButtonToggle: false
    };

    async componentDidMount() {
        const { data } = await getCategories();
        this.setState({categories : data});
    }

    handleUpdate = async () => {
        const { data } = await getCategories();
        this.setState({categories : data});
    };

    handleDelete = async categoryToDelete => {
        const originalCategories = this.state.categories;
        const categories = this.state.categories.filter(category => category.id !== categoryToDelete);
        this.setState({categories: categories});
        try {
            await deleteCategory(categoryToDelete)
        }
        catch (ex) {
            if (ex.response && ex.response.status === 404)
                toast('This category has already been deleted.');
            this.setState({categories: originalCategories});
        }
    };

    handleToggle = () => {
        let addButtonToggle = this.state.addButtonToggle;
        addButtonToggle = !addButtonToggle;
        this.setState({addButtonToggle});
    }


    render() {
        const {categories, addButtonToggle} = this.state;
        return (
            <div className="container">
                <div className="category-container">
                    <div className="add-new-category">
                        <div className="add-new-category-button">
                            <Link to="/categories/new"
                                  className="button is-link is-medium add-more-button"
                                  onClick={this.handleToggle}
                            >{addButtonToggle ? "x" : "+"}</Link>
                        </div>
                        {addButtonToggle &&
                        <Route
                            path="/categories/new"
                            render={(props) => (<CategoryForm {...props}
                                                              onUpdate={this.handleUpdate}/>)}
                        />}
                    </div>
                    {(categories.length === 0) && <h5 className="title is-5 center-text total-text">There are no categories. Please add new ones.</h5>}
                    <CategoriesList
                        categories={categories}
                        onDelete={this.handleDelete}
                    />
                </div>
            </div>
        )
    }
}


export default Categories;