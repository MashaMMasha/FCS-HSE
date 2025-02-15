using JetBrains.Annotations;
using Microsoft.VisualStudio.TestTools.UnitTesting;
namespace mini_hw1.Tests;

[TestClass]
[TestSubject(typeof(InventoryManager))]
public class InventoryManagerTest
{
    [TestMethod]
    public void Test()
    {
        InventoryManager inventoryManager = new InventoryManager(); 
        int c = inventoryManager.GetNextId();
        Assert.AreEqual(c + 1, inventoryManager.GetNextId());
    }      

}