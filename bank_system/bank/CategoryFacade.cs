namespace BankSystem;

public class CategoryFacade
{
    private Dictionary<Guid, Category> categories = new();
    public void CreateCategory(string name, OperationType type)
    {
        var category = new Category(name, type);
        categories.Add(category.id, category);
    }
    public void RenameCategory(Guid id, string newName)
    {
        categories[id].name = newName;
    }
    public void ChangeCategoryType(Guid id, OperationType newType)
    {
        categories[id].type = newType;
    }
    public void DeleteCategory(Guid id)
    {
        categories.Remove(id);
    }
}