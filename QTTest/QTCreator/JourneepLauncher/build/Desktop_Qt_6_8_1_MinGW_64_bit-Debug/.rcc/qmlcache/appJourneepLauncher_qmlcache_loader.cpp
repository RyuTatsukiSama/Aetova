#include <QtQml/qqmlprivate.h>
#include <QtCore/qdir.h>
#include <QtCore/qurl.h>
#include <QtCore/qhash.h>
#include <QtCore/qstring.h>

namespace QmlCacheGeneratedCode {
namespace _qt_qml_JourneepLauncher_main_qml { 
    extern const unsigned char qmlData[];
    extern const QQmlPrivate::AOTCompiledFunction aotBuiltFunctions[];
    const QQmlPrivate::CachedQmlUnit unit = {
        reinterpret_cast<const QV4::CompiledData::Unit*>(&qmlData), &aotBuiltFunctions[0], nullptr
    };
}
namespace _qt_qml_JourneepLauncher_UntitledProject1_Constants_qml { 
    extern const unsigned char qmlData[];
    extern const QQmlPrivate::AOTCompiledFunction aotBuiltFunctions[];
    const QQmlPrivate::CachedQmlUnit unit = {
        reinterpret_cast<const QV4::CompiledData::Unit*>(&qmlData), &aotBuiltFunctions[0], nullptr
    };
}
namespace _qt_qml_JourneepLauncher_UntitledProject1_EventListModel_qml { 
    extern const unsigned char qmlData[];
    extern const QQmlPrivate::AOTCompiledFunction aotBuiltFunctions[];
    const QQmlPrivate::CachedQmlUnit unit = {
        reinterpret_cast<const QV4::CompiledData::Unit*>(&qmlData), &aotBuiltFunctions[0], nullptr
    };
}
namespace _qt_qml_JourneepLauncher_UntitledProject1_EventListSimulator_qml { 
    extern const unsigned char qmlData[];
    extern const QQmlPrivate::AOTCompiledFunction aotBuiltFunctions[];
    const QQmlPrivate::CachedQmlUnit unit = {
        reinterpret_cast<const QV4::CompiledData::Unit*>(&qmlData), &aotBuiltFunctions[0], nullptr
    };
}
namespace _qt_qml_JourneepLauncher_UntitledProject1Content_App_qml { 
    extern const unsigned char qmlData[];
    extern const QQmlPrivate::AOTCompiledFunction aotBuiltFunctions[];
    const QQmlPrivate::CachedQmlUnit unit = {
        reinterpret_cast<const QV4::CompiledData::Unit*>(&qmlData), &aotBuiltFunctions[0], nullptr
    };
}
namespace _qt_qml_JourneepLauncher_UntitledProject1Content_Screen01_ui_0x2e_qml { 
    extern const unsigned char qmlData[];
    extern const QQmlPrivate::AOTCompiledFunction aotBuiltFunctions[];
    const QQmlPrivate::CachedQmlUnit unit = {
        reinterpret_cast<const QV4::CompiledData::Unit*>(&qmlData), &aotBuiltFunctions[0], nullptr
    };
}

}
namespace {
struct Registry {
    Registry();
    ~Registry();
    QHash<QString, const QQmlPrivate::CachedQmlUnit*> resourcePathToCachedUnit;
    static const QQmlPrivate::CachedQmlUnit *lookupCachedUnit(const QUrl &url);
};

Q_GLOBAL_STATIC(Registry, unitRegistry)


Registry::Registry() {
    resourcePathToCachedUnit.insert(QStringLiteral("/qt/qml/JourneepLauncher/main.qml"), &QmlCacheGeneratedCode::_qt_qml_JourneepLauncher_main_qml::unit);
    resourcePathToCachedUnit.insert(QStringLiteral("/qt/qml/JourneepLauncher/UntitledProject1/Constants.qml"), &QmlCacheGeneratedCode::_qt_qml_JourneepLauncher_UntitledProject1_Constants_qml::unit);
    resourcePathToCachedUnit.insert(QStringLiteral("/qt/qml/JourneepLauncher/UntitledProject1/EventListModel.qml"), &QmlCacheGeneratedCode::_qt_qml_JourneepLauncher_UntitledProject1_EventListModel_qml::unit);
    resourcePathToCachedUnit.insert(QStringLiteral("/qt/qml/JourneepLauncher/UntitledProject1/EventListSimulator.qml"), &QmlCacheGeneratedCode::_qt_qml_JourneepLauncher_UntitledProject1_EventListSimulator_qml::unit);
    resourcePathToCachedUnit.insert(QStringLiteral("/qt/qml/JourneepLauncher/UntitledProject1Content/App.qml"), &QmlCacheGeneratedCode::_qt_qml_JourneepLauncher_UntitledProject1Content_App_qml::unit);
    resourcePathToCachedUnit.insert(QStringLiteral("/qt/qml/JourneepLauncher/UntitledProject1Content/Screen01.ui.qml"), &QmlCacheGeneratedCode::_qt_qml_JourneepLauncher_UntitledProject1Content_Screen01_ui_0x2e_qml::unit);
    QQmlPrivate::RegisterQmlUnitCacheHook registration;
    registration.structVersion = 0;
    registration.lookupCachedQmlUnit = &lookupCachedUnit;
    QQmlPrivate::qmlregister(QQmlPrivate::QmlUnitCacheHookRegistration, &registration);
}

Registry::~Registry() {
    QQmlPrivate::qmlunregister(QQmlPrivate::QmlUnitCacheHookRegistration, quintptr(&lookupCachedUnit));
}

const QQmlPrivate::CachedQmlUnit *Registry::lookupCachedUnit(const QUrl &url) {
    if (url.scheme() != QLatin1String("qrc"))
        return nullptr;
    QString resourcePath = QDir::cleanPath(url.path());
    if (resourcePath.isEmpty())
        return nullptr;
    if (!resourcePath.startsWith(QLatin1Char('/')))
        resourcePath.prepend(QLatin1Char('/'));
    return unitRegistry()->resourcePathToCachedUnit.value(resourcePath, nullptr);
}
}
int QT_MANGLE_NAMESPACE(qInitResources_qmlcache_appJourneepLauncher)() {
    ::unitRegistry();
    return 1;
}
Q_CONSTRUCTOR_FUNCTION(QT_MANGLE_NAMESPACE(qInitResources_qmlcache_appJourneepLauncher))
int QT_MANGLE_NAMESPACE(qCleanupResources_qmlcache_appJourneepLauncher)() {
    return 1;
}
