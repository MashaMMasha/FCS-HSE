using JetBrains.Annotations;
using Microsoft.VisualStudio.TestTools.UnitTesting;
namespace mini_hw1.Tests;

[TestClass]
[TestSubject(typeof(Computer))]
[TestSubject(typeof(Table))]
[TestSubject(typeof(ThingFactory))]
public class ThingFactoryTest
{
    [TestMethod]
    public void TestThingFactoryWithAdditionalInformation()
    {
        var factory = new ThingFactory(new InventoryManager());
        Computer comp = factory.CreateComputer("computer", "HSE");
        Table table = factory.CreateTable("table", "Glass");
        Assert.IsNotNull(comp);
        Assert.IsNotNull(table);
        Assert.AreEqual(1, comp.Number);
        Assert.AreEqual(2, table.Number);
        Assert.AreEqual("computer", comp.Name);
        Assert.AreEqual("table", table.Name);
        Assert.AreEqual("Glass", table.Material);
        Assert.AreEqual("HSE", comp.Brand);
        
    }
    [TestMethod]
    public void TestThingFactoryWithoutAdditionalInformation()
    {
        var factory = new ThingFactory(new InventoryManager());
        Computer comp = factory.CreateComputer("computer");
        Table table = factory.CreateTable("table");
        Assert.IsNotNull(comp);
        Assert.IsNotNull(table);
        Assert.AreEqual(1, comp.Number);
        Assert.AreEqual(2, table.Number);
        Assert.AreEqual("computer", comp.Name);
        Assert.AreEqual("table", table.Name);
        Assert.AreEqual("Wood", table.Material);
        Assert.AreEqual("Apple", comp.Brand);
        
    }
}